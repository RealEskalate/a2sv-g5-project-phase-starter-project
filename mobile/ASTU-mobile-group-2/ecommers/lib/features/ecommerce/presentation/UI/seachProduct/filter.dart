
import 'package:flutter/material.dart';
import '../../../../../core/Colors/colors.dart';
import '../../../../../core/Text_Style/text_style.dart';
import '../../../../../core/border/bordrs.dart';
import 'apply_filter.dart';
import 'price_slider.dart';

class Filter extends StatefulWidget {
  const Filter({super.key});

  @override
  State<Filter> createState() => _FilterState();
}

class _FilterState extends State<Filter> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(25),
      width: double.infinity,
      height: 300,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          TextStyles(fontColor: mainText,text: 'catagory',fontSizes: 18,),
          const SizedBox(height: 10,),
          const Bordrs(
            color:Colors.white,
            hight: 40,
            width: 366,
            ),
          const SizedBox(height: 10,),
          TextStyles(fontColor: mainText,text: 'Price',fontSizes: 16),
          const SizedBox(height: 10,),
          const PriceSlider(),
          const SizedBox(height: 20,),
          const ApplyFilter()



        ],
      ),
    );
  }
}