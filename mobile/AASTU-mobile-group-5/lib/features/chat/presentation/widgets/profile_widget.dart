import 'package:flutter/material.dart';


class ProfileWidget extends StatelessWidget {
  final bool isOnline;
  final String iconUrl;
  const ProfileWidget({super.key, required this.isOnline,required this.iconUrl});

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        CircleAvatar(
          radius: 32,
          backgroundColor: Colors.blue.shade300,
          child: CircleAvatar(
            radius: 30, 
            backgroundImage: Image.asset(
              iconUrl,
            ).image,
          ),
        ),
        if (isOnline)
          Positioned(
            bottom: 0,
            right: 0,
            child: Container(
              width: 12,
              height: 12,
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.green,
                border: Border.all(
                  color: Colors.white, 
                  width: 2,
                ),
              ),
            ),
          ),
      ],
    );
  }
}