module SiteV2 exposing (main, init, update, view)

import Browser
import File exposing (File)
import File.Select as FileSelect
import Html exposing (Html, div, input, text, button)
import Html.Attributes exposing (placeholder, value)
import Html.Events exposing (onInput, onClick)
import Task


-- MAIN

main =
    Browser.sandbox { init = init, update = update, view = view }


-- MODEL

type alias Model =
    { content : String
    , fileContent : Maybe String
    }


init : Model
init =
    { content = ""
    , fileContent = Nothing
    }


-- UPDATE

type Msg
    = Change String
    | SelectFile
    | FileSelected (Result String File)
    | FileRead (Result String String)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        Change newContent ->
            ( { model | content = newContent }, Cmd.none )

        SelectFile ->
            ( model, FileSelect.file [] FileSelected )

        FileSelected (Ok file) ->
            ( model, Task.attempt FileRead (File.toString file) )

        FileSelected (Err error) ->
            ( { model | fileContent = Just ("Erreur lors de la sélection du fichier : " ++ error) }, Cmd.none )

        FileRead (Ok content) ->
            ( { model | fileContent = Just content }, Cmd.none )

        FileRead (Err error) ->
            ( { model | fileContent = Just ("Erreur lors de la lecture du fichier : " ++ error) }, Cmd.none )


-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ input
            [ placeholder "Écrire votre dictée"
            , value model.content
            , onInput Change
            ]
            []
        , button [ onClick SelectFile ] [ text "Charger un fichier" ]
        , case model.fileContent of
            Nothing ->
                text ""

            Just content ->
                div []
                    [ text "Contenu du fichier :"
                    , div [] [ text content ]
                    ]
        ]
