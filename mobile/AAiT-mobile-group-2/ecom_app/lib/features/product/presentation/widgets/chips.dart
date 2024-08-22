import 'package:flutter/material.dart';

class Chips extends StatelessWidget {
  final int number;
  final bool selected;
  const Chips({super.key, required this.number, required this.selected});

  @override
  Widget build(BuildContext context) {
    return Chip(
      backgroundColor: selected ?  Theme.of(context).primaryColor : Colors.white,
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 15),
      label: Text(
        '$number',
        style: TextStyle(fontWeight: FontWeight.bold, color: selected ? Colors.white : Colors.black),
      ),
      elevation: 2,
      shadowColor: Colors.black,
      side: BorderSide(color: selected ? Theme.of(context).primaryColor: Colors.grey.shade100),
    );
  }
}
