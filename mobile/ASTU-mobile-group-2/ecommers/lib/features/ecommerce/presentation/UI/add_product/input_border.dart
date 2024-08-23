

import 'package:flutter/material.dart';

import 'input_field.dart';

class IinputBorder extends StatelessWidget {
  final int hight;
  final bool check;
  final String text;
  final int width;
  final String placeHolder;
  final Map<String,dynamic> data;
  const IinputBorder({
    super.key,
    
    required this.hight,
    required this.check,
    required this.text,
    required this.width,
    required this.placeHolder,
    required this.data,

  });

  @override
  Widget build(BuildContext context) {

    return Container(
              width: width.toDouble(),
              height: hight.toDouble(),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color: hight > 40? const Color.fromARGB(200, 238, 238, 238): Colors.white
              ),
              child: Center(
              
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
            // Conditionally add widgets based on the 'check' variable
                  if (check) ...[
                      const Icon(Icons.add_photo_alternate_outlined),
                      const SizedBox(height: 15),
                      const Text('upload image'),
                    ] else ...[
                      InputField(
                        text:text,
                        placeHolder:placeHolder,
                        data: data,
                        
                        ),
                    ]
                  ],
                ),
              ),
            );
  }
}

