


import 'dart:io';

bool checkInput (String name,String price,String description,File? image, int type){
  if ((name.isEmpty || price.isEmpty || description.isEmpty )){
    return false;
  }
  if((name.isEmpty || price.isEmpty || description.isEmpty || image == null) && type == 0 ){
    return false;
  } 
  try{
    final double value = double.parse(price);
    if (value <= 0){
      return false;
    }

  return true;
  }catch(e){
    return false;
  }
}