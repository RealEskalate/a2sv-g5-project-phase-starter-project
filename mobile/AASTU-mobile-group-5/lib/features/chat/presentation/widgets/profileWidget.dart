import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  const ProfileWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return CircleAvatar(
      radius: 32,
      backgroundColor:Colors.blue.shade300,
      child: CircleAvatar(
        radius:150,
        backgroundImage:  Image.asset('assets/images/Alex.png',).image,
      ),
    );
    
  }
}