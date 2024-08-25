


import 'package:flutter/material.dart';

import '../../../../core/text/text.dart';


class LoginRegisterButtons extends StatelessWidget {
  final String text;
  
  const LoginRegisterButtons({
    super.key,
    required this.text,
    });

  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
    final height = MediaQuery.of(context).size.height;
    return Container(
      width: width * 0.9,
      height: 
      height * 0.055,
      decoration: const BoxDecoration(
        color: Color(0xff3F51F3),
        borderRadius: BorderRadius.all(Radius.circular(10)),

        border: Border(
          bottom: BorderSide(
            color: Color(0xff3F51F3),
            width: 2,
          ),
        ),

      ),
    child: Center(
      child: ConStTexts(
        text: text,
        color: Colors.white,
        fontSize: 15,
        fontWeight: FontWeight.w600,
        fontFamily: 'Poppins',
      
      ),
    ),
    );
  }
}