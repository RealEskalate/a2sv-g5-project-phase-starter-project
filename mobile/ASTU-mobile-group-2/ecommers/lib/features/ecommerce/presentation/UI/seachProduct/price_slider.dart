
import 'package:flutter/material.dart';
import '../../../../../core/Colors/colors.dart';

class PriceSlider extends StatefulWidget {
  const PriceSlider({super.key});

  @override
  State<PriceSlider> createState() => _PriceSliderState();
}

class _PriceSliderState extends State<PriceSlider> {
  RangeValues values = const RangeValues(0, 0.9);
  @override
  Widget build(BuildContext context) {
    RangeLabels labels = RangeLabels(
      values.start.toString(),
      values.end.toString());
    return SliderTheme(
      data: SliderTheme.of(context,).copyWith(
        trackHeight: 11,
        thumbColor: Colors.grey,
        rangeThumbShape: const RoundRangeSliderThumbShape(
          enabledThumbRadius: 10
        )
        
      ),
      child: RangeSlider( 
        activeColor: mainColor,
      values: values,
         
      labels: labels,
      onChanged: (newValues) {  
        setState (() {
          values = newValues;
        });
      },
      ),
    );
  }
}