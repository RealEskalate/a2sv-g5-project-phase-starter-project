import 'package:application1/features/product/presentation/widgets/components/size_cards.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('should return a blue card with white text when value is true',
      (widgetTester) async {
    //arrange
    await widgetTester.pumpWidget(const MaterialApp(
      home: Scaffold(
        body: SizeCards(value: true, size: 33),
      ),
    ));

    expect(find.byType(SizeCards), findsOneWidget);
    expect(find.text('33'), findsOneWidget);
    final sizeCardContainer =
        widgetTester.widget<Container>(find.byType(Container));
    final decoration = sizeCardContainer.decoration as BoxDecoration;
    expect(decoration.color, const Color.fromRGBO(63, 81, 243, 1));
    final center = sizeCardContainer.child as Center;
    final text = center.child as Text;
    final textStyle = text.style as TextStyle;
    expect(textStyle.color, Colors.white);
  });

  testWidgets('should return a white card with dark text when value is false',
      (widgetTester) async {
    //arrange
    await widgetTester.pumpWidget(const MaterialApp(
      home: Scaffold(
        body: SizeCards(value: false, size: 33),
      ),
    ));
    final sizeCardContainer =
        widgetTester.widget<Container>(find.byType(Container));
    final decoration = sizeCardContainer.decoration as BoxDecoration;
    expect(decoration.color, Colors.white);
    final center = sizeCardContainer.child as Center;
    final text = center.child as Text;
    final textStyle = text.style as TextStyle;
    expect(textStyle.color, Colors.black);
  });
}
