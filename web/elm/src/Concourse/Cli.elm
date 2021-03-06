module Concourse.Cli exposing (Cli(..), clis, downloadUrl, iconUrl, id, label)

import Api.Endpoints as Endpoints
import Url.Builder


clis : List Cli
clis =
    [ OSX, Windows, Linux ]


type Cli
    = OSX
    | Windows
    | Linux


downloadUrl : Cli -> String
downloadUrl cli =
    let
        platformName =
            case cli of
                OSX ->
                    "darwin"

                Windows ->
                    "windows"

                Linux ->
                    "linux"
    in
    Endpoints.Cli
        |> Endpoints.toString
            [ Url.Builder.string "arch" "amd64"
            , Url.Builder.string "platform" platformName
            ]


iconUrl : Cli -> String
iconUrl cli =
    let
        imageName =
            case cli of
                OSX ->
                    "apple"

                Windows ->
                    "windows"

                Linux ->
                    "linux"
    in
    "url(/public/images/" ++ imageName ++ "-logo.svg)"


label : Cli -> String
label cli =
    let
        platformName =
            case cli of
                OSX ->
                    "OS X"

                Windows ->
                    "Windows"

                Linux ->
                    "Linux"
    in
    "Download " ++ platformName ++ " CLI"


id : Cli -> String
id cli =
    case cli of
        OSX ->
            "osx"

        Windows ->
            "windows"

        Linux ->
            "linux"
