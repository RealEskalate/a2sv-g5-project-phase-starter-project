import 'package:flutter/material.dart';

class BottomNavigationBarWidget extends StatelessWidget {
  const BottomNavigationBarWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(32.0),
      color: Colors.grey[200], // Optional: To visually differentiate the area
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          Icon(
            Icons.attach_file,
            size: 40,
          ),
          Expanded(
            flex: 4,
            child: TextField(
              onTap: () {},
              decoration: InputDecoration(
                hintText: 'Write your message',
                suffixIcon: const Icon(
                  Icons.copy,
                  size: 24.0, // Adjusted size
                ),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20.0),
                ),
              ),
            ),
          ),
          IconButton(
            onPressed: () {
              // Add your phone icon action here
            },
            icon: const Icon(
              Icons.phone,
              size: 40,
            ),
          ),
          IconButton(
            onPressed: () {
              // Add your mic icon action here
            },
            icon: const Icon(
              Icons.mic,
              size: 40,
            ),
          ),
        ],
      ),
    );
  }
}
