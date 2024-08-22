import 'package:flutter/material.dart';

class SizeContainer extends StatefulWidget {
  final int size;
  final bool isSelected;
  const SizeContainer(
      {super.key, required this.size, required this.isSelected});

  @override
  State<SizeContainer> createState() => _SizeContainerState();
}

class _SizeContainerState extends State<SizeContainer> {
  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(left: 30, bottom: 20),
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: widget.isSelected
            ? const Color.fromARGB(255, 45, 69, 206)
            : const Color.fromARGB(255, 255, 252, 252),
        boxShadow: [
          BoxShadow(
            color: const Color.fromARGB(255, 213, 210, 210).withOpacity(0.5),
            spreadRadius: 1,
            blurRadius: 7,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: Text(
        widget.size.toString(),
        style: TextStyle(
            fontSize: 20,
            color: widget.isSelected ? Colors.white : Colors.black),
      ),
    );
  }
}
