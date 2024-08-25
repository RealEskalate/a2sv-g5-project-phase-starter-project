import 'package:flutter/material.dart';

class PlusButton extends StatelessWidget {
  const PlusButton({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 19, // Increase the size to accommodate the border
      height: 19,
      decoration: BoxDecoration(
        shape: BoxShape.circle,
        border: Border.all(
          color: Colors.black, // Border color
          width: 2, // Border width
        ),
      ),
      child: Container(
        width: 15,
        height: 15,
        decoration: const BoxDecoration(
          color: Color.fromARGB(255, 241, 243, 244),
          shape: BoxShape.circle,
        ),
        child: const Icon(Icons.add,
            size: 10, color: Color.fromARGB(255, 9, 6, 6)),
      ),
    );
  }
}
