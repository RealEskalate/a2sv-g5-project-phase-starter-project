import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_events.dart';
import 'package:ecommerce_app/features/product/presentation/widgets/product_list_displayer.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockGetAllProductUseCase mockGetAllProductUseCase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late MockDeleteProductUseCase mockDeleteProductUseCase;
  late MockInsertProductUseCase mockInsertProductUseCase;
  late MockGetProductUseCase mockGetProductUseCase;
  late ProductBloc productBloc;

  setUp(() {
    mockGetProductUseCase = MockGetProductUseCase();
    mockGetAllProductUseCase = MockGetAllProductUseCase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockInsertProductUseCase = MockInsertProductUseCase();
    mockDeleteProductUseCase = MockDeleteProductUseCase();
    productBloc = ProductBloc(
      getAllProductUseCase: mockGetAllProductUseCase,
      deleteProductUseCase: mockDeleteProductUseCase,
      getProductUseCase: mockGetProductUseCase,
      insertProductUseCase: mockInsertProductUseCase,
      updateProductUsecase: mockUpdateProductUsecase,
    );
  });

  testWidgets('Should display gesture with keys', (widgetTester) async {
    await widgetTester.pumpWidget(
      BlocProvider(
        create: (context) => productBloc,
        child: const MaterialApp(
          home: Scaffold(
            body: Column(
              children: [
                ProductListDisplayer(),
              ],
            ),
          ),
        ),
      ),
    );

    //debugDumpApp();
    when(mockGetAllProductUseCase.execute()).thenAnswer((_) async {
      return Right(TestingDatas.productEntityList);
    });
    debugPrint(productBloc.state.toString());
    await widgetTester.pump();
    productBloc.add(LoadAllProductEvents());
    await widgetTester.pump(const Duration(seconds: 4));
    productBloc.add(LoadAllProductEvents());
    await widgetTester.pump();
    await widgetTester.pump(const Duration(seconds: 4));

    debugPrint(productBloc.state.toString());
  });
}
