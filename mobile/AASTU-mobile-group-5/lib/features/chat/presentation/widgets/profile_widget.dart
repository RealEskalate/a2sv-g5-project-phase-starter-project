import 'package:flutter/material.dart';

class ProfileWidget extends StatelessWidget {
  final String iconUrl;
  const ProfileWidget({super.key, required this.iconUrl});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(0),
      child: CircleAvatar(
        radius: 25,
        backgroundColor: Colors.blue.shade300,
        child: CircleAvatar(
          radius: 21.5,
          backgroundImage: Image.asset(
            iconUrl,
          ).image,
        ),
      ),
    );
  }
}
