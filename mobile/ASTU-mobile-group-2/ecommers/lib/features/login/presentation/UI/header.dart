
import 'package:flutter/material.dart';

import '../../../../core/Colors/colors.dart';
import '../../../../core/text/text.dart';

class Header extends StatelessWidget {
  const Header({super.key});

  @override
  Widget build(BuildContext context) {
    final height = MediaQuery.of(context).size.height;
    final width = MediaQuery.of(context).size.width;
    return Column(
      
      children: [
        Container(
          padding: const EdgeInsets.only(bottom: 5, top: 5),
          alignment: Alignment.center,
          width: width * 0.366,
          height: height * 0.065,

          decoration:  BoxDecoration(
            borderRadius: BorderRadius.circular(10),
            color: const Color.fromARGB(255, 255, 255, 255),
            boxShadow: const [
              BoxShadow(
                color: Color(0xff3F51F3),
                blurRadius:0,
                spreadRadius: 1,
                offset: Offset(0, 0),
                blurStyle: BlurStyle.outer,

              ),
            ],
            border: Border.all(
              color: const Color(0xff3F51F3),
              width: 0.93,
            ),
          ),
          child: const Align(
            alignment: Alignment.center,
            child: ConStTexts(
              text: 'ECOM',
              color:  Color(0xff3F51F3),
              fontSize: 48,
              fontFamily: 'CaveatBrush',
              fontWeight: FontWeight.w500,
              
            ),
          ),
        ),
        const SizedBox(
          height: 70,
        ),

        const SizedBox(
          height: 35,
          child:  ConStTexts(
            text: 'Sign into your account',
            color: mainColor,
            fontSize: 26.72,
            fontWeight: FontWeight.bold,
            fontFamily: 'Poppins',
          
          ),
        ),
          
        
      ],
    );
  }
}
