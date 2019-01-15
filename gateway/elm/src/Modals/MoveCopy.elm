module Modals.MoveCopy exposing
    ( Model
    , Msg
    , newCopyModel
    , newMoveModel
    , show
    , subscriptions
    , update
    , view
    )

import Bootstrap.Alert as Alert
import Bootstrap.Button as Button
import Bootstrap.Form.Input as Input
import Bootstrap.Form.InputGroup as InputGroup
import Bootstrap.Grid as Grid
import Bootstrap.Grid.Col as Col
import Bootstrap.Grid.Row as Row
import Bootstrap.Modal as Modal
import Bootstrap.Progress as Progress
import Bootstrap.Table as Table
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Commands
import File
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode as D
import Json.Encode as E
import List
import Url
import Util


type State
    = Ready (List String)
    | Loading
    | Fail String


type alias Model =
    { state : State
    , action : Type
    , destPath : String
    , sourcePath : String
    , filter : String
    , modal : Modal.Visibility
    , alert : Alert.Visibility
    }


type Msg
    = DoAction
    | DirChosen String
    | SearchInput String
    | ModalShow String
    | GotAllDirsResponse (Result Http.Error (List String))
    | GotActionResponse (Result Http.Error String)
    | AnimateModal Modal.Visibility
    | AlertMsg Alert.Visibility
    | ModalClose



-- INIT


type Type
    = Move
    | Copy


typeToString : Type -> String
typeToString typ =
    case typ of
        Move ->
            "Move"

        Copy ->
            "Copy"


newMoveModel : Model
newMoveModel =
    { state = Loading
    , modal = Modal.hidden
    , action = Move
    , destPath = ""
    , sourcePath = ""
    , filter = ""
    , alert = Alert.shown
    }


newCopyModel : Model
newCopyModel =
    { state = Loading
    , modal = Modal.hidden
    , action = Copy
    , destPath = ""
    , sourcePath = ""
    , filter = ""
    , alert = Alert.shown
    }



-- UPDATE


fixPath : String -> String
fixPath path =
    if path == "/" then
        "Home"

    else
        String.join "/" (Util.splitPath path)


filterInvalidTargets : String -> String -> Bool
filterInvalidTargets sourcePath path =
    (path /= Util.dirname sourcePath)
        && not (String.startsWith path sourcePath)


fixAllDirResponse : Model -> List String -> List String
fixAllDirResponse model paths =
    List.filter (filterInvalidTargets model.sourcePath) paths
        |> List.map fixPath


filterCaseless : String -> String -> Bool
filterCaseless filter path =
    String.contains filter (String.toLower path)


filterAllDirs : String -> List String -> List String
filterAllDirs filter dirs =
    let
        lowerFilter =
            String.toLower filter
    in
    List.filter (String.contains lowerFilter) dirs


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        DoAction ->
            -- TODO: Implement move/copy
            case model.action of
                Move ->
                    ( model
                    , Commands.doMove GotActionResponse model.sourcePath model.destPath
                    )

                Copy ->
                    ( model
                    , Commands.doCopy GotActionResponse model.sourcePath model.destPath
                    )

        DirChosen path ->
            ( { model | destPath = path }, Cmd.none )

        SearchInput filter ->
            ( { model | filter = filter }, Cmd.none )

        GotAllDirsResponse result ->
            case result of
                Ok dirs ->
                    -- New list model means also new checked entries.
                    ( { model | state = Ready (fixAllDirResponse model dirs) }, Cmd.none )

                Err err ->
                    ( { model | state = Fail <| Util.httpErrorToString err }, Cmd.none )

        GotActionResponse result ->
            case result of
                Ok _ ->
                    -- New list model means also new checked entries.
                    ( { model | modal = Modal.hidden }, Cmd.none )

                Err err ->
                    -- TODO: Display error somehow. Change model?
                    ( model, Cmd.none )

        AnimateModal visibility ->
            ( { model | modal = visibility }, Cmd.none )

        ModalShow sourcePath ->
            ( { model
                | modal = Modal.shown
                , sourcePath = sourcePath
                , destPath = ""
                , state = Loading
              }
            , Commands.doListAllDirs GotAllDirsResponse
            )

        ModalClose ->
            ( { model | modal = Modal.hidden }, Cmd.none )

        AlertMsg vis ->
            ( { model | alert = vis }, Cmd.none )



-- VIEW


viewDirEntry : Model -> String -> Table.Row Msg
viewDirEntry model path =
    Table.tr []
        [ Table.td
            [ Table.cellAttr <| onClick (DirChosen path) ]
            [ span [ class "fas fa-lg fa-folder text-xs-right file-list-icon" ] [] ]
        , Table.td
            [ Table.cellAttr <| onClick (DirChosen path) ]
            [ text path ]
        ]


viewDirList : Model -> List String -> Html Msg
viewDirList model dirs =
    Table.table
        { options = [ Table.hover ]
        , thead =
            Table.simpleThead
                [ Table.th [ Table.cellAttr (style "width" "10%") ] []
                , Table.th [ Table.cellAttr (style "width" "90%") ] []
                ]
        , tbody =
            Table.tbody [] (List.map (viewDirEntry model) (filterAllDirs model.filter dirs))
        }


viewSearchBox : Model -> Html Msg
viewSearchBox model =
    InputGroup.config
        (InputGroup.text
            [ Input.placeholder "Filter directory list"
            , Input.attrs
                [ onInput SearchInput
                , value model.filter
                ]
            ]
        )
        |> InputGroup.successors
            [ InputGroup.span [ class "input-group-addon" ]
                [ button [] [ span [ class "fas fa-search fa-xs input-group-addon" ] [] ]
                ]
            ]
        |> InputGroup.attrs [ class "stylish-input-group input-group" ]
        |> InputGroup.view


viewContent : Model -> List (Grid.Column Msg)
viewContent model =
    [ Grid.col [ Col.xs12 ]
        [ case model.state of
            Ready dirs ->
                div []
                    [ viewSearchBox model
                    , viewDirList model dirs
                    ]

            Loading ->
                text "Loading."

            Fail message ->
                Util.buildAlert
                    model.alert
                    AlertMsg
                    Alert.danger
                    "Oh no!"
                    ("Could not rename path: " ++ message)
        ]
    ]


view : Model -> Html Msg
view model =
    Modal.config ModalClose
        |> Modal.large
        |> Modal.withAnimation AnimateModal
        |> Modal.h5 []
            [ text (typeToString model.action ++ " ")
            , span [ class "text-muted" ]
                [ text (Util.basename model.sourcePath) ]
            , if String.length model.destPath > 0 then
                span []
                    [ text " into "
                    , span [ class "text-muted" ] [ text model.destPath ]
                    ]

              else
                text " into ..."
            ]
        |> Modal.body []
            [ Grid.containerFluid []
                [ Grid.row
                    [ Row.attrs [ class "scrollable-modal-row" ] ]
                    (viewContent model)
                ]
            ]
        |> Modal.footer []
            [ Button.button
                [ Button.primary
                , Button.attrs
                    [ onClick DoAction
                    , type_ "submit"
                    , disabled
                        (String.length model.destPath
                            == 0
                            || (case model.state of
                                    Fail _ ->
                                        True

                                    _ ->
                                        False
                               )
                        )
                    ]
                ]
                [ text (typeToString model.action) ]
            , Button.button
                [ Button.outlinePrimary
                , Button.attrs [ onClick ModalClose ]
                ]
                [ text "Cancel" ]
            ]
        |> Modal.view model.modal


show : String -> Msg
show sourcePath =
    ModalShow sourcePath



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ Modal.subscriptions model.modal AnimateModal
        , Alert.subscriptions model.alert AlertMsg
        ]
