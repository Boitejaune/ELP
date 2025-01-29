module Read exposing (..)
import Parser exposing(..)


type Commande = Forward Int | Right Int| Left Int | Repeat Int (List Commande)


comright: Parser Commande
comright=succeed Right 
        |. symbol "Right"
        |. spaces
        |= int
        |. spaces
        


comforward: Parser Commande
comforward=succeed Forward 
        |. symbol "Forward"
        |. spaces
        |= int
        |. spaces
        


comleft: Parser Commande
comleft=succeed Left 
        |. symbol "Left"
        |. spaces
        |= int
        |. spaces
        

repeat: Parser Commande
repeat=succeed Repeat 
    |. symbol "Repeat"
    |. spaces
    |= int
    |. spaces
    |=lazy(\_->listinstruc)
        

listinstruc: Parser (List Commande)
listinstruc=sequence 
        { start = "["
        , separator = ","
        , end = "]"
        , spaces = spaces
        , item = read
        , trailing = Optional
        }




read: Parser Commande
read = oneOf[comforward,comleft,comright,repeat]
