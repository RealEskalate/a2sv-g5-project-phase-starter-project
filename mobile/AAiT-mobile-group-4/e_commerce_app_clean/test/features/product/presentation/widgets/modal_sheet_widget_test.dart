import 'dart:io';

import 'package:application1/features/product/presentation/widgets/components/modal_sheet_widget.dart';
import 'package:application1/features/product/presentation/widgets/components/styles/custom_button.dart';
import 'package:application1/features/product/presentation/widgets/components/styles/range_wrapper_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  setUp(() {
    HttpOverrides.global = null;
  });

  testWidgets(
      'should have a category,price, range slider and a button',
      (widgetTester) async {
    //arrange
    await widgetTester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body:  ModalSheetComponent(),
        ),
      )
    );

    expect(find.byType(Rangewrapperwidget), findsOneWidget);

    expect(find.text('Category'), findsOneWidget);
    expect(find.text('Price'), findsOneWidget);

    expect(find.byType(CustomButton), findsOneWidget);
  });
}
