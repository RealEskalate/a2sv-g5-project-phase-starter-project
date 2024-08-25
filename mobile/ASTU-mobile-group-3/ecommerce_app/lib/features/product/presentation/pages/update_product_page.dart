import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../../../core/validator/validator.dart';
import '../bloc/cubit/input_validation_cubit.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_events.dart';
import '../bloc/product_states.dart';
import '../widgets/loading_dialog.dart';
import '../widgets/product_widgets.dart';

// ignore: must_be_immutable
class UpdateProductPage extends StatelessWidget with AppBars {
  final TextEditingController nameControl = TextEditingController();
  final TextEditingController catagoryControl = TextEditingController();
  final TextEditingController priceControl = TextEditingController();
  final TextEditingController descControl = TextEditingController();
  late InputValidationCubit myCubit;
  String? id;
  static const String routes = '/update_product_page';

  UpdateProductPage({super.key});
  @override
  Widget build(BuildContext context) {
    myCubit = BlocProvider.of<InputValidationCubit>(context);
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: normalAppBar('Add Product', () {
        if (id != null) {
          BlocProvider.of<ProductBloc>(context).add(
            GetSingleProductEvents(
              id: id!,
            ),
          );
        }

        Navigator.pop(context);
      }),
      body: SafeArea(
        child: SingleChildScrollView(
          child: BlocListener<ProductBloc, ProductStates>(
            listener: (context, state) {
              if (state is SuccessfullState) {
                if (id != null) {
                  BlocProvider.of<ProductBloc>(context)
                      .add(GetSingleProductEvents(id: id!));
                }

                Navigator.pop(context);
              } else if (state is ErrorState) {
                Navigator.pop(context);
                ScaffoldMessenger.of(context)
                    .showSnackBar(SnackBar(content: Text(state.message)));
              } else if (state is LoadedSingleProductState) {
                id = state.productEntity.id;
              }
            },
            child: BlocBuilder<ProductBloc, ProductStates>(
              builder: (context, state) {
                if (state is LoadedSingleProductState) {
                  nameControl.text = state.productEntity.name;
                  BlocProvider.of<InputValidationCubit>(context).checkChanges(
                      [InputDataValidator.name, state.productEntity.name]);
                  priceControl.text = state.productEntity.price.toString();
                  BlocProvider.of<InputValidationCubit>(context).checkChanges([
                    InputDataValidator.price,
                    state.productEntity.price.toString()
                  ]);
                  descControl.text = state.productEntity.description;
                }
                return Column(
                  children: [
                    CostumInput(
                      hint: '',
                      control: nameControl,
                      text: 'Name',
                    ),
                    CostumInput(
                      hint: '',
                      control: priceControl,
                      text: 'Price',
                    ),
                    CostumInput(
                      hint: '',
                      control: descControl,
                      text: 'Description',
                      maxLine: 5,
                    ),
                    Row(
                      children: [
                        Expanded(
                          child: FillCustomButton(
                              press: () {
                                final state =
                                    BlocProvider.of<InputValidationCubit>(
                                            context)
                                        .state;

                                if (state is InputValidatedState) {
                                  if (state.isValidForUpdate()) {
                                    if (id != null) {
                                      showDialog(
                                          context: context,
                                          builder: (context) {
                                            return const LoadingDialog();
                                          });
                                      BlocProvider.of<ProductBloc>(context).add(
                                        UpdateProductEvent(
                                          id: id!,
                                          name: nameControl.text.trim(),
                                          description: descControl.text.trim(),
                                          price: priceControl.text.trim(),
                                        ),
                                      );
                                      return;
                                    } else {
                                      ScaffoldMessenger.of(context)
                                          .showSnackBar(
                                        const SnackBar(
                                          content: Text('Unknow Error Input'),
                                          duration: Duration(seconds: 2),
                                        ),
                                      );
                                    }
                                  }
                                } else {
                                  ScaffoldMessenger.of(context).showSnackBar(
                                    const SnackBar(
                                      content: Text('Invalid Input'),
                                      duration: Duration(seconds: 2),
                                    ),
                                  );
                                }
                              },
                              label: 'UPDATE'),
                        )
                      ],
                    ),
                  ],
                );
              },
            ),
          ),
        ),
      ),
    );
  }
}
