import 'package:application1/features/product/presentation/widgets/components/styles/custom_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('should have a category,price, range slider and a button',
      (widgetTester) async {
    //arrange
    bool ispressed = false;
    await widgetTester.pumpWidget(MaterialApp(
      home: Scaffold(
        body: CustomButton(
          pressed: () {
            ispressed = true;
          },
          name: 'UPDATE',
          width: 152,
          height: 50,
          fgcolor: Colors.white,
          bgcolor: const Color.fromRGBO(63, 81, 243, 1),
        ),
      ),
    ));

    expect(find.byType(CustomButton), findsOneWidget);
    expect(find.text('UPDATE'), findsOneWidget);

    await widgetTester.tap(find.byType(CustomButton));
    expect(ispressed, isTrue);
  });
}
