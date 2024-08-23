import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';

import '../../../authentication/domain/entity/user.dart';

class UserAvater extends StatelessWidget {
  final String image;
  final bool online;
  final bool story;
  
   UserAvater({super.key, required this.image, this.online = false, this.story = false});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Stack(
          children: [
            Container(
              width: 60,
              height: 60,
              decoration: BoxDecoration(
                color: Colors.amber[400],
                borderRadius: BorderRadius.circular(60),
                border: story? Border.all(color: Colors.white, width: 2, style: BorderStyle.solid, strokeAlign: BorderSide.strokeAlignCenter): Border()
              ),
              child: Image.asset(image),
            ),
            Positioned(
              bottom: 2,
              right: 5,
              child: Container(
                width: 10,
                height: 10,
                decoration: BoxDecoration(
                    color: online ? Color(0xFF0FE16D) : Color(0xFF9A9E9C),
                    borderRadius: BorderRadius.circular(10)),
              ),
            )
          ],
          
        ),
        story? const Padding(
          padding:  EdgeInsets.all(8.0),
          child:  Text("User Name", style: TextStyle(color: Colors.white, fontWeight: FontWeight.w500, fontSize: 14),),
        ): const Text("",style: TextStyle( fontSize: 0)),
      ],
    );
  }
}
