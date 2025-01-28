module Svg_test exposing (..)

import Browser
import Html exposing (Html)
import Svg exposing (..)
import Svg.Attributes exposing (..)

-- Model: Représente un point dans l'espace 2D
type alias Point =
    { x : Float, y : Float }

{-
-- Liste des segments du carré
squarePath : List (Point, Point)
squarePath =
    [ ( { x = 100, y = 100 }, { x = 200, y = 100 } ) -- Ligne du haut
    , ( { x = 200, y = 100 }, { x = 200, y = 200 } ) -- Ligne de droite
    , ( { x = 200, y = 200 }, { x = 100, y = 200 } ) -- Ligne du bas
    , ( { x = 100, y = 200 }, { x = 100, y = 100 } ) -- Ligne de gauche
    ]
-}


-- Convertir une ligne en SVG
lineToSvg : (Point, Point) -> Svg msg
lineToSvg (start, end) =
    line
        [ x1 (String.fromFloat start.x)
        , y1 (String.fromFloat start.y)
        , x2 (String.fromFloat end.x)
        , y2 (String.fromFloat end.y)
        , stroke "black"
        , strokeWidth "2"
        ]
        []

-- Vue principale : dessiner le carré
view : Html msg
view =
    svg
        [ width "300"
        , height "300"
        , viewBox "0 0 300 300"
        ]
        (List.map lineToSvg squarePath)

Path : Instructions -> Svg msg
Path instructions =
    let
        initialPosition = { x = 0, y = 0, angle = 0 }
        (svgElements, _) = drawInstruction instructions initialPosition []

-- Point d'entrée
main : Program () () msg
main =
    Browser.sandbox { init = (), update = \_ model -> model, view = \_ -> view }
