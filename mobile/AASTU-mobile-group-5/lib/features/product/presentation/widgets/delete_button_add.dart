import 'package:flutter/material.dart';


class DeleteButtonAdd extends StatelessWidget {

  const DeleteButtonAdd({super.key});

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: () {
        Navigator.pop(context);
      },
      style: OutlinedButton.styleFrom(
        side: const BorderSide(color: Colors.red),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10),
        ),
      ),
      child: const Text(
        'DELETE',
        style: TextStyle(color: Colors.red),
      ),
    );
  }
}
