module Svg_maker exposing (..)

import Svg exposing (..)
import Svg.Attributes exposing (..)
import Read exposing (..)

-- Define a position and direction type
type alias Position =
    { x : Float, y : Float, angle : Float }

-- Function to process a single instruction and generate SVG lines
drawInstruction : Instruction -> Position -> List (Svg msg) -> (List (Svg msg), Position)
drawInstruction instruction position acc =
    case instruction of
        Direction Forward steps next ->
            let
                -- Calculer la nouvelle position
                dx = toFloat steps * cos (degrees position.angle)
                dy = toFloat steps * sin (degrees position.angle)
                newX = position.x + dx
                newY = position.y + dy

                -- Créer une ligne pour ce mouvement
                lineElement =
                    line
                        [ x1 (String.fromFloat position.x)
                        , y1 (String.fromFloat position.y)
                        , x2 (String.fromFloat newX)
                        , y2 (String.fromFloat newY)
                        , stroke "black"
                        , strokeWidth "2"
                        ]
                        []
            in
            drawInstruction next { x = newX, y = newY, angle = position.angle } (lineElement :: acc)

        Direction Right angle next ->
            -- Changer l'angle pour tourner à droite
            drawInstruction next { position | angle = position.angle - toFloat angle } acc

        Direction Left angle next ->
            -- Changer l'angle pour tourner à gauche
            drawInstruction next { position | angle = position.angle + toFloat angle } acc

        Direction Repeat count subInstructions ->
            let
                repeatResult =
                    List.foldl
                        (\_ (subAcc, subPos) ->
                            drawInstruction subInstructions subPos subAcc
                        )
                        (acc, position)
                        (List.range 1 count)
            in
            drawInstruction subInstructions (Tuple.second repeatResult) (Tuple.first repeatResult)

        End ->
            (acc, position)


-- Main function to convert instructions to an SVG
draw : Instruction -> Svg msg
draw instructions =
    let
        initialPosition = { x = 0, y = 0, angle = 0 }
        (svgElements, _) = drawInstruction instructions initialPosition []
    in
    svg
        [ width "500", height "500", viewBox "-250 -250 500 500", fill "none" ]
        (List.reverse svgElements)

-- Instructions pour dessiner un carré
squareInstruction : Instruction
squareInstruction =
    Direction Repeat 4 (                     -- Répète 4 fois les instructions suivantes :
        Direction Forward 100 (              -- Avance de 100
            Direction Right 90 End           -- Tourne à droite de 90°
        )
    )

{-
    Direction Forward 100 (                -- Avance de 100
        Direction Right 90 (               -- Tourne à droite de 90°
            Direction Forward 100 (        -- Avance de 100
                Direction Right 90 (       -- Tourne à droite de 90°
                    Direction Forward 100 ( -- Avance de 100
                        Direction Right 90 ( -- Tourne à droite de 90°
                            Direction Forward 100 End -- Avance de 100 et termine
                        )
                    )
                )
            )
        )
    )
-}
-- Dessiner les instructions avec la fonction `draw`
main : Svg msg
main =
    draw squareInstruction
