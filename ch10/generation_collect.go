///main package has examples shown
// in Go Data Structures and algorithms book
package main


func GenerationCollect(){


   var currentGeneration int

   currentGeneration = 3

   var objects *[]object

   objects = GetObjectsFromOldGeneration(3)

   var object *object

   for _, object = range objects {

   var markedAlready bool

   markedAlready = IfMarked(object)
   if markedAlready {

        map[object] = true

   }
   }


}
