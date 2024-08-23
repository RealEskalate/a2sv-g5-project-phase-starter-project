

import 'package:flutter/material.dart';

class Descriptions extends StatelessWidget {
  final String text;
  const Descriptions({
    super.key,
    required this.text
    });

  @override
  Widget build(BuildContext context) {
    return Text(
      text
    );
  }
}