module Read exposing (..)
import Parser exposing(..)

type Commande = Forward Int | Right Int| Left Int | Repeat Int (List Commande)

comright: Parser Commande
comright = succeed Right 
        |. spaces
        |. symbol "Right"
        |. spaces
        |= int
        |. spaces

comforward: Parser Commande
comforward = succeed Forward 
        |. spaces
        |. symbol "Forward"
        |. spaces
        |= int
        |. spaces

comleft: Parser Commande
comleft = succeed Left 
        |. spaces
        |. symbol "Left"
        |. spaces
        |= int
        |. spaces

repeat: Parser Commande
repeat = succeed Repeat 
    |. spaces
    |. symbol "Repeat"
    |. spaces
    |= int
    |. spaces
    |= lazy(\_ -> listinstruc)

listinstruc: Parser (List Commande)
listinstruc = sequence 
        { start = "["
        , separator = ","
        , end = "]"
        , spaces = spaces
        , item = oneOf [ comforward, comleft, comright, repeat ]
        , trailing = Optional
        }

read: Parser (List Commande)
read = listinstruc