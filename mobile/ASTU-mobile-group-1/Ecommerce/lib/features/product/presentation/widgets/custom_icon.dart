import 'package:flutter/material.dart';

class CustomIconContainer extends StatelessWidget {
  final IconData icon;

  const CustomIconContainer(
    this.icon, {
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 40,
      height: 40,
      decoration: BoxDecoration(
          border: Border.all(
            color: Colors.grey,
          ),
          borderRadius: const BorderRadius.all(
            Radius.circular(8),
          )),
      child: Icon(
        icon,
        color: Colors.grey,
      ),
    );
  }
}
