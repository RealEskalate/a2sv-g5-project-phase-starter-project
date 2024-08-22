import 'package:flutter/material.dart';
import 'custom_outlined_button.dart';

class ModalSheet extends StatelessWidget {
  const ModalSheet({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(left: 20.0,right: 20,top: 20),
      child: SizedBox(
        height: 300,
        width: double.infinity,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text('Category'),
            const SizedBox(
              height: 8,
            ),
            const SizedBox(
              height: 45,
              child: TextField(
                decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.black12))),
              ),
            ),
            const SizedBox(
              height: 25,
            ),
            const Text('Price'),
            RangeSlider(
                activeColor: Theme.of(context).primaryColor,
                values: const RangeValues(0.2, 0.8),
                onChanged: (RangeValues rv) {}),
            const SizedBox(
              height: 50,
            ),
            CustomOutlinedButton(
                backgroundColor: Theme.of(context).primaryColor,
                foregroundColor: Colors.white,
                borderColor: Theme.of(context).primaryColor,
                buttonWidth: double.maxFinite,
                buttonHeight: 45,
                child: const Text(
                                'APPLY',
                                style: const TextStyle(
                                    fontWeight: FontWeight.w600),
                              ),
                onPressed: () {}),
          ],
        ),
      ),
    );
  }
}
