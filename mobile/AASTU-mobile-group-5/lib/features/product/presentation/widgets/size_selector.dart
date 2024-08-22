import 'package:flutter/material.dart';

class SizeSelector extends StatefulWidget {
  const SizeSelector({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _SizeSelectorState createState() => _SizeSelectorState();
}

class _SizeSelectorState extends State<SizeSelector> {
  int _selectedSize = 37;

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      scrollDirection: Axis.horizontal,
      child: Row(
        children: [38, 39, 40, 41, 42, 43, 44, 45].map((size) {
          return Padding(
            padding: const EdgeInsets.all(3.0),
            child: ChoiceChip(
              backgroundColor: Colors.white,
              shadowColor: Colors.red,
              label: Text(
                '$size',
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                  color: _selectedSize == size ? Colors.white : Colors.black,
                ),
              ),
              selected: _selectedSize == size,
              showCheckmark: false,
              selectedColor: const Color.fromARGB(255, 54, 104, 255),
              onSelected: (selected) {
                setState(() {
                  _selectedSize = size;
                });
              },
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(16),
                side: const BorderSide(color: Colors.transparent),
              ),
              labelPadding:
                  const EdgeInsets.symmetric(vertical: 8.0, horizontal: 8.0),
            ),
          );
        }).toList(),
      ),
    );
  }
}
