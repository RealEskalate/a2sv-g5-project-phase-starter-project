import 'package:flutter/material.dart';

SnackBar snackBar(String message){
  return SnackBar(
    content: Text(message) ,
    backgroundColor: Colors.green,
  );
}

SnackBar errorsnackBar(String message){
  return SnackBar(
    content: Text(message) ,
    backgroundColor: Colors.red,
  );
}
