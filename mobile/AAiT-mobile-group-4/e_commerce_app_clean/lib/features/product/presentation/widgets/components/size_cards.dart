import 'package:flutter/material.dart';

class SizeCards extends StatelessWidget {
  final bool value;
  final int size;
  const SizeCards({super.key, required this.value, required this.size});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Container(
        width: 60,
        height: 60,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(8),
          color: value ? const Color.fromRGBO(63, 81, 243, 1) : Colors.white,
        ),
        child: Center(
          child: Text(
            '$size',
            style: TextStyle(
              fontFamily: 'Poppins',
              fontWeight: FontWeight.w500,
              fontSize: 20,
              color: value ? Colors.white : Colors.black,
            ),
          ),
        ),
      ),
    );
  }
}
