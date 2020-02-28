# RID1000-A proprietary protocol reader 


Read arguments from RID1000-A genset panel + RS485 + MOXA NPort 5150, and returns a json object.

Version of RID1000-A  - 18.04.2016 SH RID1000-A_Vers.1.0.29М

Programm flags:

-ip - MOXA ip address (defaut value "localhost");

-port - request TCP port (defaut value 2001);

`--=Important Notice=--`

`Not all parameter names are defined.`

`Undefined parameters go in the format "R021229"`

Example:

`rid_prop -ip=10.10.10.10 -port=2002 `
