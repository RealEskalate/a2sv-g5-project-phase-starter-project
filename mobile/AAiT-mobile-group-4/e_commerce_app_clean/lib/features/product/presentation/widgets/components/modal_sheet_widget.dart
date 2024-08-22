import 'package:flutter/material.dart';

import 'styles/custom_button.dart';
import 'styles/range_wrapper_widget.dart';
import 'styles/text_style.dart';

class ModalSheetComponent extends StatefulWidget {
  const ModalSheetComponent({super.key});

  @override
  State<ModalSheetComponent> createState() => _ModalSheetComponentState();
}

class _ModalSheetComponentState extends State<ModalSheetComponent> {
  @override
  Widget build(BuildContext context) {
    return Container(
      height: 338,
      padding: const EdgeInsets.all(32),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const CustomTextStyle(name: 'Category',weight: FontWeight.w500,size: 16),
          const SizedBox(height: 10),
          Container(
            width: double.infinity,
            height: 40,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(8),
              border: Border.all(
                  color:
                      const Color.fromRGBO(217, 217, 217, 1)),
            ),
          ),
          const SizedBox(height: 16),
          const CustomTextStyle(name: 'Price',weight: FontWeight.w500,size: 16),
          const SizedBox(height: 10),
          const Rangewrapperwidget(),
          const SizedBox(height: 20,),
          CustomButton(
            width: double.infinity,
            height: 44,
            name: 'APPLY',
            textBgColor: Colors.white,
            fgcolor: Theme.of(context).primaryColor,
            bgcolor: Theme.of(context).primaryColor,
          )
        ],
      ),
    );
  }
}