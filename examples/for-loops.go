package main

func main() {
  // this also demostrates scope
  var i int

  for
    // i is shadowed because we did not use := to declare and assign
    //  instead we only assigned using =
    i = 0
    i < 2
    i++ { continue }

  // this is the same as above but we declare a new var i before assigning
  for i:=0; i<2; i++ { continue }

  // while loop
  for i <= 0 { i-- }

  // infinite while loop
  for { break }
}
