import 'package:flutter/material.dart';

class UpdateButton extends StatelessWidget {
  final dynamic product;

  const UpdateButton({super.key, required this.product,});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {
       Navigator.pushNamed(context, '/update', arguments: product);
      },
      style: ElevatedButton.styleFrom(
        backgroundColor: const Color.fromARGB(255, 54, 104, 255),
        foregroundColor: Colors.white,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(12),
        ),
      ),
      child: const Text('UPDATE'),
    );
  }
}
