import 'package:flutter/material.dart';

class Redirect extends StatelessWidget {
  final String text;
  final String buttonText;
  final Widget navigateTo;

  const Redirect({
    super.key,
    required this.text,
    required this.buttonText,
    required this.navigateTo,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 0 ,vertical: 10),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            text,
            style: const TextStyle(color: Colors.grey),
          ),
          TextButton(
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => navigateTo),
              );
            },
            child: Text(
              buttonText,
              style: const TextStyle(color:Color(0xFF3F51F3)),
            ),
          ),
        ],
      ),
    );
  }
}