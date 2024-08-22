import 'package:flutter/material.dart';

class AddButtonWidget extends StatelessWidget {
  const AddButtonWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return FloatingActionButton(
      onPressed: () {
        Navigator.pushNamed(context, '/add_update');
      },
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(30),
      ),
      backgroundColor: const Color.fromARGB(255, 54, 104, 255),
      child: const Icon(
        Icons.add,
        size: 35,
        color: Colors.white,
      ),
    );
  }
}
