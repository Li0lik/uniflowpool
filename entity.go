package uniflowpool

//Entity
/*
   Structure of any object
   **Value keep object
   **Previous this is link to previous object
   **Next this is link to next object
*/
type Entity struct {
	Previous interface{}
	Next     interface{}

	Value interface{}
}
