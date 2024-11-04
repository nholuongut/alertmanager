{-
   Alertmanager API
   API of the Prometheus Alertmanager (https://github.com/nholuongut/alertmanager)

   OpenAPI spec version: 0.0.1

   NOTE: This file is auto generated by the openapi-generator.
   https://github.com/openapitools/openapi-generator.git
   Do not edit this file manually.
-}


module Data.InlineResponse200 exposing (InlineResponse200, decoder, encoder)

import Dict exposing (Dict)
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (optional, required)
import Json.Encode as Encode


type alias InlineResponse200 =
    { silenceID : Maybe String
    }


decoder : Decoder InlineResponse200
decoder =
    Decode.succeed InlineResponse200
        |> optional "silenceID" (Decode.nullable Decode.string) Nothing


encoder : InlineResponse200 -> Encode.Value
encoder model =
    Encode.object
        [ ( "silenceID", Maybe.withDefault Encode.null (Maybe.map Encode.string model.silenceID) )
        ]
