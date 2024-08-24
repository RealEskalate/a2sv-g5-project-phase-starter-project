import 'package:flutter/material.dart';


class CircularIndicator extends StatelessWidget {
  const CircularIndicator({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        color: Colors.white,
        height: 100,
        child: const Center(
          child: CircularProgressIndicator(
            color: (Color(0xff800080)),
          ),
        ),
      ),
    );
  }
}
