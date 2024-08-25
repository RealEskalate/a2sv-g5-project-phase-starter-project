import 'package:flutter/material.dart';

class StoryAvatar extends StatelessWidget {
  const StoryAvatar({super.key});

  @override
  Widget build(BuildContext context) {
    return Stack(children: [
      Column(
        children: [
          Stack(
            children: [
              Container(
                padding: const EdgeInsets.all(3),
                width: 70,
                height: 70,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.blue,
                  border: Border.all(
                    color: Colors.white,
                    width: 2,
                  ),
                ),
                child: const CircleAvatar(
                  radius: 30,
                  backgroundImage: AssetImage('assets/images/user.png'),
                ),
              ),
            ],
          ),
          const SizedBox(height: 10),
          const Text(
            'My Status',
            style: TextStyle(color: Colors.white),
          ),
        ],
      )
    ]);
  }
}
