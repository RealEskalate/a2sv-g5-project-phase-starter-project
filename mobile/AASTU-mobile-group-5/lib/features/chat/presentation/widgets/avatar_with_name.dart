import 'package:flutter/material.dart';
import 'profile_widget.dart';

// ignore: camel_case_types
class avatar_with_name extends StatelessWidget {
  final String name;
  const avatar_with_name({
    super.key,
    required this.name,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(
        right: 20,
      ),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const ProfileWidget(isOnline:true, iconUrl: 'images/leather_shoe_1.png',),
          Text(
            name,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }
}
