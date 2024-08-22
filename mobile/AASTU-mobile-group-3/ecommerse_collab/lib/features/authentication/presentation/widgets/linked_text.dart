import 'package:flutter/material.dart';

class LinkedText extends StatelessWidget {
  final String description;
  final String linkedText;
  final Widget navigateTo;

  const LinkedText({super.key, this.description = "", required this.linkedText, required this.navigateTo});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(20),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        
        children: [
          Text(description, 
          style: const TextStyle(
      
            fontFamily: 'Poppins',
            fontWeight: FontWeight.w400,
            fontSize: 16,
            color:  Color.fromARGB(255, 172, 161, 161)
      
      
            
          ),),
          GestureDetector(
            onTap: () {
              Navigator.push(context, MaterialPageRoute(builder: (BuildContext context) => navigateTo));
            },
            child: Text(linkedText, 
            
            style: const TextStyle(
              fontFamily: 'Poppins',
              fontWeight: FontWeight.w400,
              fontSize: 16,
              color: Color(0xff3F51F3)),),
          ),
        ],
      ),
    );
  }
}