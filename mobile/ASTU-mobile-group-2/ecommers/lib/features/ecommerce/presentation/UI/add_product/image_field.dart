

import 'dart:io';

import 'package:flutter/material.dart';


class ImageField extends StatelessWidget {
  final int hight;
  final bool check;
  final String text;
  final int width;
  final String imageUrl;
  final File? localImage;
 
  const ImageField({
    super.key,
    
    required this.hight,
    required this.check,
    required this.text,
    required this.width,
    this.imageUrl = '',
    this.localImage,
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
              child: localImage != null ?Image.file(localImage!, height: hight.toDouble(),
                    fit: BoxFit.fill,
                    ): imageUrl == ''?const Column(
               
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                          // Conditionally add widgets based on the 'check' variabl
                       
                   
                
                    const Icon(Icons.add_photo_alternate_outlined),
                    
                    const SizedBox(height: 15),
                    const Text('upload image'),
                    
                    
                
                ],
                
              ):Image.network(imageUrl, height: hight.toDouble(),
                    fit: BoxFit.fill,),
            );
  }
}

