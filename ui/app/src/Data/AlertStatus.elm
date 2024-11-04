{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/nholuongut/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.AlertStatus exposing (AlertStatus, State(..), decoder, encoder)

import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias AlertStatus =
    { state : State
    , silencedBy : List String
    , inhibitedBy : List String
    , mutedBy : List String
    }


type State
    = Unprocessed
    | Active
    | Suppressed


decoder : Decoder AlertStatus
decoder =
    Decode.succeed AlertStatus
        |> required "state" stateDecoder
        |> required "silencedBy" (Decode.list Decode.string)
        |> required "inhibitedBy" (Decode.list Decode.string)
        |> required "mutedBy" (Decode.list Decode.string)


encoder : AlertStatus -> Encode.Value
encoder model =
    Encode.object
        [ ( "state", stateEncoder model.state )
        , ( "silencedBy", Encode.list Encode.string model.silencedBy )
        , ( "inhibitedBy", Encode.list Encode.string model.inhibitedBy )
        , ( "mutedBy", Encode.list Encode.string model.mutedBy )
        ]


stateDecoder : Decoder State
stateDecoder =
    Decode.string
        |> Decode.andThen
            (\str ->
                case str of
                    "unprocessed" ->
                        Decode.succeed Unprocessed

                    "active" ->
                        Decode.succeed Active

                    "suppressed" ->
                        Decode.succeed Suppressed

                    other ->
                        Decode.fail <| "Unknown type: " ++ other
            )


stateEncoder : State -> Encode.Value
stateEncoder model =
    case model of
        Unprocessed ->
            Encode.string "unprocessed"

        Active ->
            Encode.string "active"

        Suppressed ->
            Encode.string "suppressed"