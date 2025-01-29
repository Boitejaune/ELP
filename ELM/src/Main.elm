module Main exposing (main, init, update, view)

import Browser
import Html exposing (Html, div, input, text, p)
import Html.Attributes exposing (placeholder, value)
import Html.Events exposing (onInput)
import Svg exposing (Svg)
import Svg_maker exposing (draw)
import Read exposing (read)
import Parser exposing (run)


-- MAIN

main =
    Browser.sandbox { init = init, update = update, view = view }


-- MODEL

type alias Model msg =
    { content : String
    , svgOutput : Maybe (Svg msg)
    }


init : Model msg
init =
    { content = ""
    , svgOutput = Nothing }


-- UPDATE

type Msg
    = Change String


update : Msg -> Model msg -> Model msg
update msg model =
    case msg of
        Change newContent ->
            let
                parsedInstruction =
                    case run read newContent of
                        Ok instruction ->
                            Just (draw [instruction])

                        Err _ ->
                            Nothing
            in
            { model | content = newContent, svgOutput = parsedInstruction }


-- VIEW

view : Model msg -> Html Msg
view model =
    div [Html.Attributes.class "page"]
        [ div [Html.Attributes.class "field"][text "Notez votre code TcTurtle içi:"]
        ,
        div [Html.Attributes.class "field"]
            [input
            [ placeholder "example: [Repeat 4 [Forward 100, Right 90]]"
            , value model.content
            , onInput Change
            , Html.Attributes.class "field__input"
            ]
            []
            ]
        , div []
            [case model.svgOutput of
            Just svg -> Html.map (\_ -> Change "") svg -- Transforme `Svg msg` en `Html Msg`
            Nothing -> p [Html.Attributes.class "error"] [text "Invalid input or no drawing yet."]]
        ]
