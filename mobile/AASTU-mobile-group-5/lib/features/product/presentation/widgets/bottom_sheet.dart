import 'package:flutter/material.dart';
import 'range_slidder.dart';

class FilterBottomSheet extends StatelessWidget {
  const FilterBottomSheet({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Text(
                'Category',
                style: TextStyle(fontSize: 20),
              ),
            ],
          ),
          const SizedBox(
            height: 5,
          ),
          TextField(
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
          const SizedBox(height: 16),
          const SizedBox(child: MyRange()),
          const SizedBox(height: 16),
          ElevatedButton(
            onPressed: () {
              // Add your apply logic here
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: const Color.fromARGB(255, 54, 104, 255),
              foregroundColor: Colors.white,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(12),
              ),
            ),
            child: const Text('APPLY'),
          ),
        ],
      ),
    );
  }
}
