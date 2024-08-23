import 'package:flutter/material.dart';

class Rangewrapperwidget extends StatefulWidget {
  const Rangewrapperwidget({super.key});

  @override
  State<Rangewrapperwidget> createState() => _RangewrapperwidgetState();
}

class _RangewrapperwidgetState extends State<Rangewrapperwidget> {
  
  RangeValues _currentRangeValues = const RangeValues(300, 700);
  
  @override
  Widget build(BuildContext context) {
    return RangeSlider(
      values: _currentRangeValues,
      min: 0,
      max: 1000,
      divisions: null, 
      onChanged: (RangeValues values) {
        setState(() {
          _currentRangeValues = values;
        });
      },
      activeColor: Colors.blue[800],
    
    );
  }
}