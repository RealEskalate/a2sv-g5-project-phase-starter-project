import 'package:ecom_app/features/product/presentation/widgets/modal_sheet.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main(){

  testWidgets('Modal Sheet Widget Test', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(const MaterialApp(
      home: Scaffold(
        body: ModalSheet(),
      ),
    ));

    expect(find.byType(ModalSheet), findsOneWidget);
  });
}