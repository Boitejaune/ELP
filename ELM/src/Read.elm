module Read exposing (..)
import Parser exposing(..)


type Commande = Forward | Right | Left | Repeat
type Instruction = Direction Commande Int (Instruction ) | End

keyword : Parser Commande
keyword =
    oneOf
        [  succeed Forward |. symbol "Forward"
        , succeed Right |. symbol "Right"
        , succeed Left |. symbol "Left"
        , succeed Repeat |. symbol "Repeat"
        ]

read: Parser Instruction
read = 
    succeed Direction
        |.symbol "["
        |. spaces
        |= keyword
        |. spaces
        |= int
        |. spaces
        |= oneOf[lazy (\_ -> read),succeed End]
        |. spaces
        |.symbol "]"
    
example : String -> Result (List DeadEnd) Instruction
example input =
    Parser.run read input


{-
type Boolean
  = MyTrue
  | MyFalse
  | MyOr Boolean Boolean

boolean : Parser Boolean
boolean =
  oneOf
    [ succeed MyTrue
        |. keyword "true"
    , succeed MyFalse
        |. keyword "false"
    , succeed MyOr
        |. symbol "("
        |. spaces
        |= lazy (\_ -> boolean)
        |. spaces
        |. symbol "||"
        |. spaces
        |= lazy (\_ -> boolean)
        |. spaces
        |. symbol ")"
    ]
-}