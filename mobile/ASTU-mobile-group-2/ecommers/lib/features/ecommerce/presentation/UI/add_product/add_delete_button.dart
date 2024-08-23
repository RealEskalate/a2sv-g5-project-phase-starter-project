

import 'package:flutter/material.dart';

import '../../../../../core/const/width_height.dart';

class AddDeleteButton extends StatelessWidget {
  final Color color;
  final String text;
  final Color borderColor;
  const AddDeleteButton({
    super.key,
    required this.color,
    required this.text,
    required this.borderColor
    });

  @override
  Widget build(BuildContext context) {
    double width = WidthHeight.screenWidth(context);
    double height = WidthHeight.screenHeight(context);
    return Container(
              width: width*0.85,
              height: 0.053*height,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color : color,
                border: Border(
                  top: BorderSide(color: borderColor),
                  left: BorderSide(color: borderColor),
                  right: BorderSide(color: borderColor),
                  bottom: BorderSide(color: borderColor),
                
                )
              ),
              child: Center(
              
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
            // Conditionally add widgets based on the 'check' variable
                        if (text != 'DELETE') ...[
                        Text(text,style: const TextStyle(color: Colors.white),),
                        ] else ...[
                          Text(text,style: const TextStyle(color: Color.fromARGB(255, 255, 1, 1)),),
                        ]
                  ],
                ),
              ),
            );
  }
}