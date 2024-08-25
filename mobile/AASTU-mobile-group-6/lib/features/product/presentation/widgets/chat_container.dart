import 'package:flutter/material.dart';

class SingleText extends StatelessWidget{
  final String chat;
  final String parentColor;
  SingleText({required this.chat, this.parentColor = 'blue'});

  @override
  Widget build(BuildContext context){
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(20),
        // color: Color.fromRGBO(63, 81, 243, 1),
      ),
      child: Padding(
        padding: const EdgeInsets.only(left: 5, right: 3),
child: Text(
  chat,
  style: (this.parentColor == 'blue')
      ? TextStyle(color: Colors.white)
      : TextStyle(color: Colors.black),
),        
      )
      
    );
  }
}