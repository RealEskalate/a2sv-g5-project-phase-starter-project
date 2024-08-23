import 'dart:io';

import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/presentation/bloc/product_bloc.dart';
import 'package:application1/features/product/presentation/pages/update_page.dart';
import 'package:application1/features/product/presentation/widgets/components/styles/text_field_styles.dart';
import 'package:bloc_test/bloc_test.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';

class MockProductBloc extends MockBloc<ProductEvent, ProductState>
    implements ProductBloc {}

class MockFile extends Mock implements File {}

class FakeProductEntity extends Fake implements ProductEntity {}

void main() {
  late MockProductBloc mockProductBloc;
  late ProductEntity mockProduct;
  
  setUp(() {
    mockProductBloc = MockProductBloc();
    mockProduct = const ProductEntity(
      id: '1',
      name: 'Test Product',
      description: 'Test Description',
      price: 50.0,
      imageUrl: 'path/to/image',
    );
    
    when(() => mockProductBloc.state).thenReturn(ProductInitial());
  });

  Widget buildTestableWidget() {
    return MaterialApp(
      home: BlocProvider<ProductBloc>(
        create: (context) => mockProductBloc,
        child: UpdatePage(selectedProduct: mockProduct),
      ),
    );
  }

  testWidgets('renders UpdatePage with initial data',
      (WidgetTester tester) async {
    await tester.pumpWidget(buildTestableWidget());

    expect(find.text('Update Product'), findsOneWidget);
    expect(find.text(mockProduct.name), findsOneWidget);
    expect(find.text(mockProduct.description), findsOneWidget);
  });

  testWidgets('displays selected image when an image is picked',
      (WidgetTester tester) async {
    await tester.pumpWidget(buildTestableWidget());

    final Finder uploadImageIcon = find.byIcon(Icons.image_outlined);

    expect(uploadImageIcon, findsOneWidget);

    await tester.tap(uploadImageIcon);
    await tester.pumpAndSettle();

    expect(find.byType(Image), findsOneWidget);
  });

  testWidgets('triggers ProductBloc event when update button is pressed',
      (WidgetTester tester) async {
    await tester.pumpWidget(buildTestableWidget());

    await tester.enterText(find.byType(CustomTextField).first, 'Updated Name');
    await tester.enterText(find.byType(CustomTextField).last, '60.0');

    await tester.tap(find.text('UPDATE'));
    await tester.pumpAndSettle();

    verify(() => mockProductBloc.add(any())).called(1);
  });

  testWidgets('shows success snackbar and navigates when update is successful',
      (WidgetTester tester) async {
    whenListen(mockProductBloc,
        Stream.fromIterable([ProductUpdatedState(mockProduct)]));

    await tester.pumpWidget(buildTestableWidget());

    expect(find.text('Successfully Updated Product'), findsOneWidget);
    expect(find.byType(UpdatePage), findsNothing);
  });

  testWidgets('shows error snackbar when update fails',
      (WidgetTester tester) async {
    whenListen(mockProductBloc,
        Stream.fromIterable([const ProductErrorState( 'error')]));

    await tester.pumpWidget(buildTestableWidget());

    expect(find.text('error'), findsOneWidget);
  });
}
