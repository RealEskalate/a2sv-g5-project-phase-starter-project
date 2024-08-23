

import 'package:flutter/material.dart';

import '../../../../core/text/text.dart';


class RegistrationAppBar extends StatelessWidget {
  const RegistrationAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    final height = MediaQuery.of(context).size.height;
    final width = MediaQuery.of(context).size.width;
    return Container(
          padding: const EdgeInsets.only(bottom: 5, top: 5),
          alignment: Alignment.center,
          width: width * 0.166,
          height: height * 0.035,

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
              fontSize: 16,
              fontFamily: 'CaveatBrush',
              fontWeight: FontWeight.w900,
              
            ),
          ),
        );
  }
}