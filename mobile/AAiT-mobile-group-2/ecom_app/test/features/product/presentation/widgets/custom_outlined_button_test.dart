import 'package:ecom_app/features/product/presentation/widgets/custom_outlined_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('CustomOutlinedButton widget test', (WidgetTester tester) async {
    bool wasPressed = false;

    await tester.pumpWidget(MaterialApp(
      home: Scaffold(
        body: CustomOutlinedButton(
          backgroundColor: Colors.blue,
          foregroundColor: Colors.white,
          borderColor: Colors.red,
          buttonWidth: 200,
          buttonHeight: 50,
          onPressed: () {
            wasPressed = true;
          },
          child: const Text('Test Button'),
        ),
      ),
    ));

    expect(find.byType(CustomOutlinedButton), findsOneWidget);

    expect(find.text('Test Button'), findsOneWidget);

    final outlinedButton =
        tester.widget<OutlinedButton>(find.byType(OutlinedButton));
    expect(outlinedButton.style?.fixedSize?.resolve({}), const Size(200, 50));

    expect(outlinedButton.style?.backgroundColor?.resolve({}), Colors.blue);

    expect(outlinedButton.style?.foregroundColor?.resolve({}), Colors.white);

    final borderSide = outlinedButton.style?.side?.resolve({});
    expect(borderSide?.color, Colors.red);
    expect(borderSide?.width, 1.0);

    await tester.tap(find.byType(CustomOutlinedButton));
    expect(wasPressed, isTrue);
  });
}
