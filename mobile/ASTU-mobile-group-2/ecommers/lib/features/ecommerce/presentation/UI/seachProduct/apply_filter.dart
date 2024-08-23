

import 'package:flutter/material.dart';

import '../../../../../core/Colors/colors.dart';
import '../../../../../core/const/width_height.dart';

class ApplyFilter extends StatelessWidget {
  final String text;

  const ApplyFilter({
    super.key,
    this.text = 'Apply'
    
    });

  @override
  Widget build(BuildContext context) {
    double width = WidthHeight.screenWidth(context);
    double height = WidthHeight.screenHeight(context);
    return Container(
              width: 0.85*width,
              height: height*0.053,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color : const Color.fromARGB(255, 95, 110, 245),
                border: const Border(
                  top: BorderSide(color: mainColor),
                  left: BorderSide(color: mainColor),
                  right: BorderSide(color: mainColor),
                  bottom: BorderSide(color: mainColor),
                
                )
              ),
              child:  Center(
              
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
            // Conditionally add widgets based on the 'check' variable
                      
                        Text(text,style: const TextStyle(color: Colors.white),),
                  ]
                       
                ),
              ),
            );
  }
}