module Svg_maker exposing (..)

import Svg exposing (..)
import Svg.Attributes exposing (..)
import Read exposing (..)

-- Define a position and direction type
type alias Position =
    { x : Float, y : Float, angle : Float }

-- Function to process a single instruction and generate SVG lines
drawInstruction : Commande -> Position -> List (Svg msg) -> (List (Svg msg), Position)
drawInstruction instruction position acc =
    case instruction of
        Forward steps ->
            let
                dx = toFloat steps * cos (degrees position.angle)
                dy = toFloat steps * sin (degrees position.angle)
                newX = position.x + dx
                newY = position.y + dy

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
            (lineElement :: acc, { x = newX, y = newY, angle = position.angle })

        Right angle ->
            (acc, { position | angle = position.angle + toFloat angle })

        Left angle ->
            (acc, { position | angle = position.angle - toFloat angle })

        Repeat count commands ->
            let
                repeatResults =
                    List.foldl
                        (\_ (subAcc, subPos) ->
                            List.foldl
                                (\cmd (cmdAcc, cmdPos) ->
                                    drawInstruction cmd cmdPos cmdAcc
                                )
                                (subAcc, subPos)
                                commands
                        )
                        (acc, position)
                        (List.range 1 count)
            in
            repeatResults



-- Main function to convert instructions to an SVG
draw : List Commande -> Svg msg
draw instructions =
    let
        initialPosition = { x = 0, y = 0, angle = 0 }
        (svgElements, _) =
            List.foldl
                (\cmd (acc, pos) -> drawInstruction cmd pos acc)
                ([], initialPosition)
                instructions
    in
    svg
        [ width "500", height "500", viewBox "-250 -250 500 500", fill "none" ]
        (List.reverse svgElements)