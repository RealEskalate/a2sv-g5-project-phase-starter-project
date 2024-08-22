import 'package:flutter/material.dart';

class SearchBar extends StatelessWidget {
  final TextEditingController controller;
  final VoidCallback onFilterPressed;

  const SearchBar({
    super.key,
    required this.controller,
    required this.onFilterPressed,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Expanded(
          child: TextField(
            controller: controller,
            decoration: InputDecoration(
              suffixIcon: const Icon(Icons.arrow_forward),
              hintText: 'Leather',
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(6),
              ),
            ),
          ),
        ),
        const SizedBox(width: 10),
        Container(
          padding: const EdgeInsets.all(8),
          width: 56,
          height: 56,
          decoration: BoxDecoration(
            color: const Color.fromARGB(255, 54, 104, 255),
            borderRadius: BorderRadius.circular(6),
          ),
          child: IconButton(
            onPressed: onFilterPressed,
            icon: const Icon(
              Icons.filter_list,
              color: Colors.white,
            ),
          ),
        ),
      ],
    );
  }
}
